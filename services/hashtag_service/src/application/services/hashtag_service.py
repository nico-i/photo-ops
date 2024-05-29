import logging

import torch
import transformers
from grpc import RpcContext

from shared.python.__generated__.proto.services.hashtag_service.v1.hashtag_service_pb2 import (
    GetHashtagsRequest, GetHashtagsResponse)
from shared.python.__generated__.proto.services.hashtag_service.v1.hashtag_service_pb2_grpc import \
    HashtagServiceServicer
from shared.python.infrastructure.logging.logger import get_logger


class HashtagService(HashtagServiceServicer):
    def __init__(self, model_id: str = "meta-llama/Meta-Llama-3-8B-Instruct"):
        self.__pipeline = transformers.pipeline(
            "text-generation",
            model=model_id,
            model_kwargs={"torch_dtype": torch.bfloat16},
            device_map="auto",
            )
        self.__logger = get_logger()
        self.__system_msg = {"role": "system", "content": "You are an Instagram hashtag generator that only \
                             responds with a string in the form of a csv. The string must contain a list of a \
                             maximum of 30 hashtags, which are related to the given image description. \
                             Choose popular Instagram hashtags that are relevant to the given description \
                             and will maximize the reach of the final post. Do not include any special characters or quotations."}
        
    def get_hashtags(self, request: GetHashtagsRequest, context: RpcContext) -> GetHashtagsResponse:
        self.__logger.info(f"request: {request}")
       
        messages = [
            self.__system_msg,
            {"role": "user", "content": request.caption},
        ]
        
        prompt = self.__pipeline.tokenizer.apply_chat_template(
                messages, 
                tokenize=False, 
                add_generation_prompt=True
        )

        terminators = [
            self.__pipeline.tokenizer.eos_token_id,
            self.__pipeline.tokenizer.convert_tokens_to_ids("<|eot_id|>")
        ]

        outputs = self.__pipeline(
            prompt,
            max_new_tokens=256,
            pad_token_id = self.__pipeline.tokenizer.eos_token_id,
            eos_token_id=terminators,
            do_sample=True,
            temperature=0.6,
            top_p=0.9,
        )
        
        prediction = outputs[0]["generated_text"][len(prompt):]
        
        res = GetHashtagsResponse(hashtags_csv=prediction)
        
        self.__logger.info(f"response: {res}")
        
        return res
        