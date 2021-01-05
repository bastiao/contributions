# Contributions: a few scripts for Phabricator/CI integration

The idea is to keep the communication between phabricator and jenkins in your command line, such as check pending differentials or run a build job in Jenkins.


## List differentials

This is an example to list and watch for differential contributions. 

The `watch` is an optional flag. 


```
$ contributions arc --list --watch 
```

Result: List of open differentials.

```

⭐ Starting pha-go with arc command.
         List:  false
         Watch:  true
         Params:  

🚒 Looking for the contributions for today. 
📃 Endpoint:  https://phabricator.localdomain
⌛ Token:  cli-hash

🚒 Watching. 


🎆 Open or pending differentials:
        🐊 URI:  https://phabricator.localdomain/D100
        Branch:  imp/example
        StatusName:  Accepted
        Repo:  rREPO01
```


## Build branch over CI

Here, you can easily start a new build manually, by command-line. 

It is easier to get the identifier queue address.

```
$ contributions jenkins --branch=branch_example_param --params-ci="linux_agent=linux04"
```

The result will be look for Jenkins and start a new job with a few parameters:


```
⭐ Starting pha-go with jenkins command.
         List:  
         Watch:  
         Params:  code_branch=test/1.0.0,linux_agent=linux04
         Revision:  0

🏃 Jenkins mode.

🙅 Jenkins Nodes:

        📗 Node is online master
        📗 Node is online windows09
        📗 Node is online linux04
        📗 Node is online linux20

🎃 Latest job:


         - Last Success Build: [] 
         - Duration:  912 seconds

🎃 Current build:


         - Params:  map[test/1.0.0,linux_agent=linux04]
         📕 Jenkins Build Id:  6720
         - Job:  &{0xc000122a00 0xc000013560 /job/Pipeline}
         - Building Number:  774
         - Params:  [{linux_agent=linux04} {BranchDevops */master}]
         - Duration:  0 seconds
         - Running:  true
         - Output:
```


## How to run for development? 

There are only a few examples to make life easier for development: 

### CI/Jenkins

Run a specific branch in the Jenkins

```
PHA_ARGS="jenkins --branch=branch_example_param --params-ci=linux_agent=linux04 --repo=example01 --revision=100" make run
```
### Differentals 

Run in development mode: 

```
PHA_ARGS="arc" make run
```


### Documentations 

This allow to check a list of studies with a specific keyword to match in the title.
For instance, it will look for pages with keyword "Support" and stop on find the date different than "2020", and check only the titles that match P1 or P2.

./bin/contributions docs --list --query="Support" --filter="2020" --match "P1|P2"