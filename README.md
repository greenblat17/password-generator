
# Password Generator

Telegram bot that generates a password according to certain patterns


## Features

- Easy to remember password
- Each service has a different password
- Difficulty of cracking a password
- If someone found out your password from service1, they never found out your password from service2


## Documentation

User information to create a password:
- secret key
- secret number
- registration service

Password generation algorithm:
- the password looks like: service_secret-key_nums[3]
- service: the first and last letter of the service on which registration takes place
- secret key: your keyword
- secret number: any 4 digits
- nums:  
    1. take the first and last letters from the name of the service
    2. replace them with numbers that corresponded to them in the push-button layout of the phone
    3. take the last digit so that the sum of these 3 digits is equal to 9, and if this is impossible, then 19
    4. the resulting 3 digits are converted to a three-digit number
    5. the first two digits of the secret number are taken, the second two digits of the secret number are added, and then summed with the number from step 4
## Optimizations

A refactor is needed for the main file. 


## Screenshots

![App Screenshot](https://sun9-63.userapi.com/impg/KbMcEHz3Vysbq6leefYiz-RxcfpNktkeItQW2A/s_sffpALMEk.jpg?size=840x660&quality=96&sign=b32404a32c783eda40091ba01dc5f59a&type=album)

