https://www.cise.ufl.edu/~sahni/dsaaj/enrich/c16/tries.htm

https://www.cse.iitd.ac.in/~mausam/courses/col106/spring2021/lectures/12-tries.pdf

The standard trie for a set of strings S is an ordered tree such that:
- each node but the root is labeled with a character
- the children of a node are alphabetically ordered
- the paths from the external nodes to the root yield the strings of S

A standard trie supports the following operations on a preprocessed text in time O(m), where m = |X|
- [word matching]: find the first occurrence of word X in the text
- [prefix matching]: find the first occurrence of the longest prefix of word X in the text