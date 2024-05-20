# Interview assignment

Assignment: https://docs.google.com/document/d/1mMLxSQlPyd9q1cD0d4__fUlK_5NOgJupsxlobQvvcOM/edit

### To run
rename .env.example -> .env
and run in terminal
```
make run
```

### All commands

## Run
```
make run
```

## Test
```
make test
```

## Coverage
```
make cover
```

# TODO:
By default extraction delay is 30s. To change it to 5 min. change `fetchDelay` function in main.go

In case of scaling make sense to extract data in batches. Batches are divided between instances.
For example, first instance extracts 1-1000 Users, second 1001-2000 and so on...

