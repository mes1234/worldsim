# Domain Analysis

Domain under investigation is a system which is arranged in map like structure divided into discreet __cells__. Every __cell__ is a smallest portion and cannot be divided further. Goal of this system is to allow to perform analysis of user actions impact on system. Eg. : 

__Population__ places some __building__ in one location and what effect it has on nearby __life__ and __air__ condition.
## Interaction analysis
Map is arranged in X-Y coordinate and closest __neighbours__ are considered 8 __cells__ next to it with the same __rate of impact__. Interaction is based on domain specific __threshold__ eg:

__Air__ purification rate in some __cell__ exceeds its cell based __threshold__ so __neighbour__ cells obtain 1/8 of purification rate above __threshold__.

# Domains:
## Matter
### Definition
 Matter is something what exists everywhere and it doesn't matter its quantity eg:
 * Air
 * Surface

 It has some property or properties which can be reduced to __Quality__ eg:
  * Air quality - contamination/ PM10/2.5 content
  * Surface - fertility 

If any other domain is interested in given aspect of matter it always preffer it to be max __Quality__

### Domains:
#### Surface
#### Air

----------
## Resource
### Definition
### Domains:
#### Life
#### Minerals
-------------
## Belong
### Definition
### Domains:
#### Population
#### Buildings

 