# Closest common manager

## Overview

_Bureaucr.at_ is a typical hierarchical organisation. Claire, its CEO, has a hierarchy of employees reporting to her and 
each employee can have a list of other employees reporting to him/her. An employee with at least one report is called a 
Manager. 

We will implement a corporate directory for _Bureaucr.at_, with an interface to find the closest common Manager between 
two employees. All employees eventually report up to the CEO.

## Hypothesis

1. We will suppose that the hierarchy is strict (the graph is not cyclic), so that it cannot be that A is reporting to B 
and at the same time that B is reporting to A. 

2. The common manager of A and B, when A is reporting to B will be B, by definition. In general, the common manager of A
and C, if A is reporting to B and B is reporting to C, will be C. 

3. There will be always one and only one manager at the top, called CEO, who is not reporting to anyone. (For this time,
we don't take the stakeholders in account.)

## Technical implementation

The problem can be solved making use of the common lowest ancestor algorithm in a DAG (directed acyclic graph).

### Input 

The information of the organization is read from a csv file, where there is a listing all employees. Every record has
three fields: the id of the employee, their name and the ids of the employees who are reporting to them.

The rest of the input will be given as arguments in the command line: in first place the path to the csv file, and then
the two names of the employees for which we want to know the closest common manager. 

## How to run..

###  the tests

`go test`

###  the compiler 

`go build -o ccm.exe`

###  the program 

`ccm path_to_file employee_1 employee_2`