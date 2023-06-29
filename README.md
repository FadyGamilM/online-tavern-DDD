# online-tavern-DDD

# The business

➜ In this project I am trying to build an online tavern <br>
➜ The core-domain is the tavern itself<br>
➜ The sub-domains are [Customer], [Employee], [Financial], [Suppliers]<br>

# DDD structure

### **Entity**

➜ Entity is the placeholder for a domain/sub-domain model which is used to keep information about the data<br>
➜ Entity is identified by an ID<br>
➜ Entity is mutable and it can be modified from other domains <br>
➜ Same entity can be used in multiple sub-domains, thats why I decided to create a separate package for the `En tities` and not define my entities inside the domain package<br>

### **Value-Object**

➜ Value object is immutable, and it has no identifier, its a representation for a data and once its created it can't be modified, if you need to update the data, you must delete it and create a new value object<br>
➜ `Transaction` is a value object in my project, so I implemented its fields with lowercase because I won't allow any sub-domain to mutate it<br>

### **Aggregates**

➜ Aggregate is used to represent the concepts about my domain that can't be representeed by a single entity or value-objects, for example the `customer` is a `person` and `customer` can buy a `product` and also `customer` can perform `transaction` <br>
➜ Aggregate's data shouldn't be exposed outside the aggregate to other sub-domains, thats why I decided to keep its paramters as lowercase defined.
