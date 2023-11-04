
# Password Generator

Telegram bot that generates a password according to certain patterns


## Features

- Easy to remember password
- Each service has a different password
- Difficulty of cracking a password
- If someone found out your password from service1, they never found out your password from service2


## Documentation

#### User information to create a password
- secret key
- secret number
- registration service

#### Password generation algorithm
- the password looks like: service_secret-key_nums[3]
- service: the first and last letter of the service on which registration takes place
- secret key: your keyword
- secret number: any 4 digits
- nums:  
    1. take the first and last letters from the name of the service
    2. replace them with numbers that corresponded to them in the push-button layout of the phone
    3. take the last digit so that the sum of these 3 digits is equal to 9, and if this is impossible, then 19
    4. the resulting 3 digits are converted to a three-digit number
    5. the first two digits of the secret number are taken, the second two digits of the secret number are added.
    6. numbers from 4 and 5 points are added

#### Example
 - service: github
 - secret key: secret
 - secret number: 1234

1. First letter: G; Last letter: B
2. G = num1 = 4; B = num2 = 2
3. num1 + num2 = 4 + 2 = 6 => num3 = 9 - 6 = 3
4. num1 = 4, num2 = 2, num3 = 3 => 423
5. 1234 <=> 12 + 34 = 46
6. 423 + 46 = 469

 

## Optimizations

A refactor is needed for the main file. 


## Screenshots

![App Screenshot](https://sun9-31.userapi.com/impg/i7ya6VgSBP4cCw1L9ZlGyxEwHDbD1JOJlIXkRA/L_aApBlyMAE.jpg?size=922x774&quality=96&sign=cc9754e9e19ab605543bcf2c4066fe77&type=album)

