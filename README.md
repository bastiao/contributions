# Contributions: a few scripts for Phabricator/CI integration

The idea is to keep the communication between phabricator and jenkins in your command line, such as check pending differentials or run a build job in Jenkins.


## List differentials

```
./contributions arc --list --watch 
```

Result: List of open differentials.

## Build branch over CI

```
./contributions jenkins --branch=imp/T2123 --params-ci="agent=agent01"
```


## How to run for development? 

```
PHA_ARGS="arc" make run
```

```
PHA_ARGS="jenkins --branch=branch_example_param --params-ci=linux_agent=local01 --repo=example01 --revision=D100" make run
```