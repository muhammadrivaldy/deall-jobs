# DeallJobs Service

## Preparation

You only need to set up docker & docker-compose in your local environment.
After that you can running it with command `docker-compose up -d` for easily development process.

## How to use it?

Import this `api_contract.yaml` to the your Postman tools and all of APIs it will appear after successfully.
Try to login with email `admin@example.com` & password `SayaAdmin`. With that account, you will have all access of our APIs.
You can create another user with email (`admin@example.com`) and make sure you put `user_type_id` as `2`. Why? Because `type 2` is a role for regular user and only have an access for get user.
