# Gitea Task Checker

A gitea action (works with github as well) that checks if all the tasks in the pull request are completed. I use it for my daily-journal to make sure that I complete all the tasks in the PR for the day before merging back to main.

## Usage

### Example workflow ( for PR) :
```yaml
name: 'Test Gitea Task Checker'
on:
  pull_request_target:
    types: [opened, edited, reopened, synchronize]
jobs:
  use-go-action:
    runs-on: ubuntu-23.10
    steps:
      - name: Use Go Action  
        id: simple-action
        uses: https://github.com/gopi487krishna/gitea-taskchecker@main
        with:
          username: meow

      - name: Print Status
        run: echo '${{ steps.gitea-taskchecker.outputs.status }}'
```
Next just create a pull request with a few tasks. Whenever you check or uncheck a task,the workflow will run and fail if all the tasks in PR are not checked.


## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.


## License

[MIT](https://choosealicense.com/licenses/mit/)
