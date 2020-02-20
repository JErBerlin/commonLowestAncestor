# Representation of an organisation with graphs

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

The input will be given in a csv file, listing all employees. Every record will have three fields: it will hold the id 
of the employee, its name and the ids of the employees who are reporting to it.

### Interface

An interface will allow to enter the ids of two of the employees and will return the id of the common manager. There 
will be always a solution since there is the CEO as the top manager.



