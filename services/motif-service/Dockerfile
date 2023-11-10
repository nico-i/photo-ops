 bbox = process_image(args.input_path, args.debug_path)

    # Convert numpy.int64 elements to Python native int for JSON serialization
    bbox_dict = {
        "x": int(bbox[0]), 
        "y": int(bbox[1]), 
        "width": int(bbox[2]), 
        "height": int(bbox[3])
    }
    
    print(json.dumps(bbox_dict))