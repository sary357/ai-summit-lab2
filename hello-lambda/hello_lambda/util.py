import datetime
import random
import string

def generate_unique_id(length=12):
    """
    Generates a unique ID with the format:
    YYYYMMDDHHmmss + 10 characters (random charset + number)
    """
    # Get the current date and time
    now = datetime.datetime.now()
    
    # Format the date and time
    date_time_str = now.strftime('%Y%m%d%H%M%S')
    
    # Generate a random 10-character string
    random_chars = ''.join(random.choices(string.ascii_letters + string.digits, k=length))
    
    # Combine the date/time and random characters
    unique_id = date_time_str + "-" +random_chars
    
    return unique_id

