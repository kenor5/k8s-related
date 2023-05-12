

def handle(req):
    """handle a request to the function
    Args:
        req (str): request body
    """
    array = req.split()
    res = 0
    for i in array:
        res += int(i)
    return res
