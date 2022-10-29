# Test project for fintech trainee program

## Project requirements
- docker
- docker-compose

## Start the project
First of all you should create your personal .env file in the root directory. The list of required values is stored in the `.env_development` file, feel free to copy this values to your file. Or execute the command:
````
cp .env_development .env
````
### Setting up the project
 - start the project vie `make start_project` command
 - run initial migrations `make run_local_migrations`

## Testing

After running the project the page with swagger doc will be available via the link: http://127.0.0.1:8000/swagger/index.html

## Additional tasks

1. Unit/integration -> done
2. Allow to process big size files with limited RAM -> done in [PR #6](https://github.com/rostikts/fintech_test_project/pull/6)
3. Endpoint for returning the csv -> done in [PR #7](https://github.com/rostikts/fintech_test_project/pull/7)