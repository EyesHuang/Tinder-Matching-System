# interview


## Question 1
Check out [this](https://hackmd.io/wp_lbzWrSc-vJFEpUb4OrQ?view) golang program. What happens when this program runs?


## Question 2
You are required to implement an API that queries a user's recent 100
purchased products. The API's RTT time should be lower than 50ms, so you need to use
Redis as the data store. How would you store the data in Redis? How would you minimize
memory usage?


## Question 3
Please explain the difference between rolling upgrade and re-create
Kubernetes deployment strategies, and the relationship between rolling upgrade and
readiness probe.


## Question 4
Check out the following SQL. Of index A or B, which has better performance
and why?
```
SELECT * FROM orders WHERE user_id = ? AND created_at >= ? AND status = ?
index A : idx_user_id_status_created_at(user_id, status, created_at)
index B : idx_user_id_created_at_status(user_id, created_at, status)
index C : idx_user_id_created_at(user_id, created_at)
```

## Question 5
In the Kafka architecture design, how does kafka scale consumer-side
performance? Does its solution have any drawbacks? Is there any counterpart to this
drawback?


## Question 6
Please follow the following requirements to implement an HTTP server and post
your GitHub repo link.
Design an HTTP server for the Tinder matching system. The HTTP server must support the
following three APIs:
1. AddSinglePersonAndMatch : Add a new user to the matching system and find any
possible matches for the new user.
2. RemoveSinglePerson : Remove a user from the matching system so that the user
cannot be matched anymore.
3. QuerySinglePeople : Find the most N possible matched single people, where N is a
request parameter.
Here is the matching rule:
- A single person has four input parameters: name, height, gender, and number of
wanted dates.
- Boys can only match girls who have lower height. Conversely, girls match boys who
are taller.
- Once the girl and boy match, they both use up one date. When their number of dates
becomes zero, they should be removed from the matching system.
Note : Please do not use other databases such as MySQL or Redis, just use in-memory
data structure which in application to store your data.
Other requirements :
- Unit test
- Docker image
- Structured project layout
- API documentation
- System design documentation that also explains the time complexity of your API
- You can list TBD tasks.