# interview


## Question 1
Check out [this](https://hackmd.io/wp_lbzWrSc-vJFEpUb4OrQ?view) golang program. What happens when this program runs?

### Answer
<details>
  <summary>Click me</summary>


The code fragment has two problems.
- Array Size Too Large
- Deadlock

You can refer to the [link](https://github.com/EyesHuang/interview/q1) for fix version.


<u>Array Size Too Large</u>
It has the following error because Go has a limit on symbol size, typically around 2GB. ([ref link](https://github.com/golang/go/issues/9862))

**Error**
```
Build Error: go build -o C:\Users\YongTeng\interview\q1\__debug_bin4284365681.exe -gcflags all=-N -l .
# q1
./main.go:33:16: main..stmp_0: symbol too large (800000000000 bytes > 2000000000 bytes)
./main.go:33:16: main..stmp_1: symbol too large (800000000000 bytes > 2000000000 bytes) (exit status 1)
```

Correct `for _ = range [10e10]uint64{}` to `for i := 0; i < 10e10; i++`.


<u>Deadlock</u>
After the correction, it has the following error due to deadlcok.

**Error**
```
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0xc00000a050?)
	C:/Program Files/Go/src/runtime/sema.go:62 +0x25
sync.(*WaitGroup).Wait(0xc00000a050)
	C:/Program Files/Go/src/sync/waitgroup.go:116 +0x8b
main.main()
	C:/Users/YongTeng/interview/bitorpo/interview/q1/main.go:45 +0x270
```

Add `if else` statement for consistent lock ordering.

**Origin**
```
func transfer(from *User, to *User, amount uint64) {
	from.Lock.Lock()
	to.Lock.Lock()
	defer from.Lock.Unlock()
	defer to.Lock.Unlock()

	if from.Balance >= amount {
		from.Balance -= amount
		to.Balance += amount
	}
}
```

**Correction**
```
func transfer(from *User, to *User, amount uint64) {
	if from.ID < to.ID {
		from.Lock.Lock()
		defer from.Lock.Unlock()
		to.Lock.Lock()
		defer to.Lock.Unlock()
	} else {
		to.Lock.Lock()
		defer to.Lock.Unlock()
		from.Lock.Lock()
		defer from.Lock.Unlock()
	}

	if from.Balance >= amount {
		from.Balance -= amount
		to.Balance += amount
	}
}
```
</details>

## Question 2
You are required to implement an API that queries a user's recent 100
purchased products. The API's RTT time should be lower than 50ms, so you need to use
Redis as the data store. How would you store the data in Redis? How would you minimize
memory usage?

### Answer
<details>
  <summary>Click me</summary>

Redis Lists are a better choice for storing user purchases due to their ordered nature and efficient operations. The necessary operations (push, trim, and range) are well-supported by Redis Lists.

A list in Redis can be treated as a queue, allowing us to easily add new purchases to the top of the list with `LPUSH` and retrieve the most recent 100 purchases with `LRANGE`. While trimming the list with `LTRIM` is not required to use `LRANGE`, it helps to keep memory usage efficient by maintaining the list at a manageable size.

```
# Data structure: purchases:<user_id> <product_id>

# Add a purchase
$ redis-cli LPUSH purchases:user_1234 product_5678

# Trim the list to the latest 100 purchases
$ redis-cli LTRIM purchases:user_1234 0 99

# Get the most recent 100 purchases
$ redis-cli LRANGE purchases:user_1234 0 99
```

</details>

## Question 3
Please explain the difference between rolling upgrade and re-create
Kubernetes deployment strategies, and the relationship between rolling upgrade and readiness probe.

### Answer
<details>
  <summary>Click me</summary>

Kubernetes prioritizes high availability with rolling upgrades by default. This method uses readiness probes for a seamless transition:

- New pods are verified healthy before replacing old ones, minimizing downtime.
- Quick rollbacks are ensured if problems occur.

The simpler re-create strategy, though faster, has downsides:
- Service disruption occurs during the update.
- No gradual rollout means all pods are replaced at once, potentially causing issues.

</details>

## Question 4
Check out the following SQL. Of index A or B, which has better performance
and why?
```
SELECT * FROM orders WHERE user_id = ? AND created_at >= ? AND status = ?
```
index A : idx_user_id_status_created_at(user_id, status, created_at)
index B : idx_user_id_created_at_status(user_id, created_at, status)
index C : idx_user_id_created_at(user_id, created_at)

### Answer
For optimal query performance, consider using **Index B (idx_user_id_created_at_status)**.
This index is specifically designed to efficiently handle queries that filter by `user_id` (exact match), `created_at` (range), and `status` (exact match). The order of columns in the index (user_id, created_at, status) ensures it can be fully utilized for all these conditions, leading to faster query execution.

## Question 5
In the Kafka architecture design, how does kafka scale consumer-side
performance? Does its solution have any drawbacks? Is there any counterpart to this
drawback?

### Answer
<details>
  <summary>Click me</summary>


</details>

## Question 6
<details>
  <summary>Click me</summary>

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

</details>

### Answer
Please refer to the details under [q6](https://github.com/EyesHuang/interview/tree/main/q6) folder.
- API documentation: [/q6/docs/swagger.yaml](https://github.com/EyesHuang/interview/blob/main/q6/docs/swagger.yaml)
- System design documentation: [/q6/README.md](https://github.com/EyesHuang/interview/blob/main/q6/README.md)
