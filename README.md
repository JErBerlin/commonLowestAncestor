# Representation of an organisation with graphs

## Overview

Bureaucrat is a typical hierarchical organisation. Claire, its CEO, has a hierarchy of employees reporting to her and 
each employee can have a list of other employees reporting to him/her. An employee with at least one report is called a 
Manager. 

We will implement a corporate directory for Bureaucrat, with an interface to find the closest common Manager between 
two employees. All employees eventually report up to the CEO.

## Technical implementation

The problem can be solved making use of the common lowest ancestor algorithm in a DAG (directed acyclic graph).

### Input 

The input will be through a file, listing all employees. Every record will have three fields: it will hold the id of the
employee, its name and the ids of the employees who are reporting to it.

### Interface

An interface will allow to enter the ids of two of the employees and will return the id of the common manager. (There 
will be always a solution since there is the CEO as top manager.)  



