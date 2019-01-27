# PostgreSQL Performance Playground

## Usage

```bash
docker-compose run bench
```

## multicolumns pkg

This package tests the performance of `=` operation on two columns, one with low cardinality, one with high cardinality,
and compares the differences of different table sizes, indexes, column types and query modes.

Example result:

```bash
▶ docker-compose run bench
Creating pg-perf-playground_postgres_1 ... done
goos: linux
goarch: amd64
pkg: github.com/rueian/pg-perf-playground/multicolumns
BenchmarkMultiColumns/rows=10000/btree(c1_int,c2_int)/batch=1/mode=WhereIn-6         	    5000	    372008 ns/op
BenchmarkMultiColumns/rows=10000/btree(c1_int,c2_int)/batch=1/mode=Foreach-6         	    2000	    807023 ns/op
BenchmarkMultiColumns/rows=10000/btree(c1_int,c2_int)/batch=100/mode=WhereIn-6       	    1000	   1842017 ns/op
BenchmarkMultiColumns/rows=10000/btree(c1_int,c2_int)/batch=100/mode=Foreach-6       	      50	  29612202 ns/op
BenchmarkMultiColumns/rows=10000/btree(c1_int,c2_uuid)/batch=1/mode=WhereIn-6        	    3000	    340135 ns/op
BenchmarkMultiColumns/rows=10000/btree(c1_int,c2_uuid)/batch=1/mode=Foreach-6        	    2000	    834017 ns/op
BenchmarkMultiColumns/rows=10000/btree(c1_int,c2_uuid)/batch=100/mode=WhereIn-6      	     500	   2045498 ns/op
BenchmarkMultiColumns/rows=10000/btree(c1_int,c2_uuid)/batch=100/mode=Foreach-6      	      50	  28318702 ns/op
BenchmarkMultiColumns/rows=10000/btree(c1_str,c2_uuid)/batch=1/mode=WhereIn-6        	    5000	    372531 ns/op
BenchmarkMultiColumns/rows=10000/btree(c1_str,c2_uuid)/batch=1/mode=Foreach-6        	    2000	    813992 ns/op
BenchmarkMultiColumns/rows=10000/btree(c1_str,c2_uuid)/batch=100/mode=WhereIn-6      	    1000	   2201916 ns/op
BenchmarkMultiColumns/rows=10000/btree(c1_str,c2_uuid)/batch=100/mode=Foreach-6      	      50	  30394476 ns/op
BenchmarkMultiColumns/rows=10000/btree(c1_str,c2_str)/batch=1/mode=WhereIn-6         	    5000	    358928 ns/op
BenchmarkMultiColumns/rows=10000/btree(c1_str,c2_str)/batch=1/mode=Foreach-6         	    2000	    844446 ns/op
BenchmarkMultiColumns/rows=10000/btree(c1_str,c2_str)/batch=100/mode=WhereIn-6       	    1000	   2266788 ns/op
BenchmarkMultiColumns/rows=10000/btree(c1_str,c2_str)/batch=100/mode=Foreach-6       	      50	  26142256 ns/op
BenchmarkMultiColumns/rows=10000/btree(c2_int)/batch=1/mode=WhereIn-6                	    3000	    348397 ns/op
BenchmarkMultiColumns/rows=10000/btree(c2_int)/batch=1/mode=Foreach-6                	    3000	    796560 ns/op
BenchmarkMultiColumns/rows=10000/btree(c2_int)/batch=100/mode=WhereIn-6              	    1000	   1769686 ns/op
BenchmarkMultiColumns/rows=10000/btree(c2_int)/batch=100/mode=Foreach-6              	      50	  26140772 ns/op
BenchmarkMultiColumns/rows=10000/btree(c2_uuid)/batch=1/mode=WhereIn-6               	    3000	    364623 ns/op
BenchmarkMultiColumns/rows=10000/btree(c2_uuid)/batch=1/mode=Foreach-6               	    2000	    816110 ns/op
BenchmarkMultiColumns/rows=10000/btree(c2_uuid)/batch=100/mode=WhereIn-6             	    1000	   1894396 ns/op
BenchmarkMultiColumns/rows=10000/btree(c2_uuid)/batch=100/mode=Foreach-6             	      50	  29603114 ns/op
BenchmarkMultiColumns/rows=10000/btree(c2_str)/batch=1/mode=WhereIn-6                	    5000	    345225 ns/op
BenchmarkMultiColumns/rows=10000/btree(c2_str)/batch=1/mode=Foreach-6                	    2000	    736453 ns/op
BenchmarkMultiColumns/rows=10000/btree(c2_str)/batch=100/mode=WhereIn-6              	    1000	   1892399 ns/op
BenchmarkMultiColumns/rows=10000/btree(c2_str)/batch=100/mode=Foreach-6              	      50	  25846040 ns/op
BenchmarkMultiColumns/rows=10000/hash(c2_int)/batch=1/mode=WhereIn-6                 	    5000	    317349 ns/op
BenchmarkMultiColumns/rows=10000/hash(c2_int)/batch=1/mode=Foreach-6                 	    2000	    716147 ns/op
BenchmarkMultiColumns/rows=10000/hash(c2_int)/batch=100/mode=WhereIn-6               	    1000	   1656187 ns/op
BenchmarkMultiColumns/rows=10000/hash(c2_int)/batch=100/mode=Foreach-6               	      50	  31612118 ns/op
BenchmarkMultiColumns/rows=10000/hash(c2_uuid)/batch=1/mode=WhereIn-6                	    5000	    350609 ns/op
BenchmarkMultiColumns/rows=10000/hash(c2_uuid)/batch=1/mode=Foreach-6                	    2000	    804823 ns/op
BenchmarkMultiColumns/rows=10000/hash(c2_uuid)/batch=100/mode=WhereIn-6              	    1000	   1769744 ns/op
BenchmarkMultiColumns/rows=10000/hash(c2_uuid)/batch=100/mode=Foreach-6              	      50	  27651280 ns/op
BenchmarkMultiColumns/rows=10000/hash(c2_str)/batch=1/mode=WhereIn-6                 	    5000	    343161 ns/op
BenchmarkMultiColumns/rows=10000/hash(c2_str)/batch=1/mode=Foreach-6                 	    2000	    722210 ns/op
BenchmarkMultiColumns/rows=10000/hash(c2_str)/batch=100/mode=WhereIn-6               	    1000	   1836133 ns/op
BenchmarkMultiColumns/rows=10000/hash(c2_str)/batch=100/mode=Foreach-6               	      50	  27700760 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c1_int,c2_int)/batch=1/mode=WhereIn-6       	    5000	    356619 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c1_int,c2_int)/batch=1/mode=Foreach-6       	    2000	    851732 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c1_int,c2_int)/batch=100/mode=WhereIn-6     	    1000	   1924356 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c1_int,c2_int)/batch=100/mode=Foreach-6     	      50	  30117244 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c1_int,c2_uuid)/batch=1/mode=WhereIn-6      	    5000	    334818 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c1_int,c2_uuid)/batch=1/mode=Foreach-6      	    2000	    740919 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c1_int,c2_uuid)/batch=100/mode=WhereIn-6    	    1000	   2069214 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c1_int,c2_uuid)/batch=100/mode=Foreach-6    	      50	  29281244 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c1_str,c2_uuid)/batch=1/mode=WhereIn-6      	    5000	    364655 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c1_str,c2_uuid)/batch=1/mode=Foreach-6      	    2000	    802135 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c1_str,c2_uuid)/batch=100/mode=WhereIn-6    	    1000	   2250684 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c1_str,c2_uuid)/batch=100/mode=Foreach-6    	      50	  29032284 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c1_str,c2_str)/batch=1/mode=WhereIn-6       	    5000	    365292 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c1_str,c2_str)/batch=1/mode=Foreach-6       	    2000	    798921 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c1_str,c2_str)/batch=100/mode=WhereIn-6     	    1000	   2322749 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c1_str,c2_str)/batch=100/mode=Foreach-6     	      50	  29712294 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c2_int)/batch=1/mode=WhereIn-6              	    5000	    339665 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c2_int)/batch=1/mode=Foreach-6              	    2000	    759778 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c2_int)/batch=100/mode=WhereIn-6            	    1000	   1801297 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c2_int)/batch=100/mode=Foreach-6            	      50	  30728294 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c2_uuid)/batch=1/mode=WhereIn-6             	    5000	    365050 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c2_uuid)/batch=1/mode=Foreach-6             	    2000	    845021 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c2_uuid)/batch=100/mode=WhereIn-6           	    1000	   1968802 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c2_uuid)/batch=100/mode=Foreach-6           	      50	  29388246 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c2_str)/batch=1/mode=WhereIn-6              	    5000	    357523 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c2_str)/batch=1/mode=Foreach-6              	    2000	    771200 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c2_str)/batch=100/mode=WhereIn-6            	    1000	   2021417 ns/op
BenchmarkMultiColumns/rows=1000000/btree(c2_str)/batch=100/mode=Foreach-6            	      50	  36678522 ns/op
BenchmarkMultiColumns/rows=1000000/hash(c2_int)/batch=1/mode=WhereIn-6               	    5000	    342755 ns/op
BenchmarkMultiColumns/rows=1000000/hash(c2_int)/batch=1/mode=Foreach-6               	    2000	    758722 ns/op
BenchmarkMultiColumns/rows=1000000/hash(c2_int)/batch=100/mode=WhereIn-6             	    1000	   1756438 ns/op
BenchmarkMultiColumns/rows=1000000/hash(c2_int)/batch=100/mode=Foreach-6             	      50	  27427660 ns/op
BenchmarkMultiColumns/rows=1000000/hash(c2_uuid)/batch=1/mode=WhereIn-6              	    5000	    341668 ns/op
BenchmarkMultiColumns/rows=1000000/hash(c2_uuid)/batch=1/mode=Foreach-6              	    2000	    820019 ns/op
BenchmarkMultiColumns/rows=1000000/hash(c2_uuid)/batch=100/mode=WhereIn-6            	    1000	   1840662 ns/op
BenchmarkMultiColumns/rows=1000000/hash(c2_uuid)/batch=100/mode=Foreach-6            	      50	  31432426 ns/op
BenchmarkMultiColumns/rows=1000000/hash(c2_str)/batch=1/mode=WhereIn-6               	    5000	    335188 ns/op
BenchmarkMultiColumns/rows=1000000/hash(c2_str)/batch=1/mode=Foreach-6               	    2000	    770165 ns/op
BenchmarkMultiColumns/rows=1000000/hash(c2_str)/batch=100/mode=WhereIn-6             	    1000	   1846859 ns/op
BenchmarkMultiColumns/rows=1000000/hash(c2_str)/batch=100/mode=Foreach-6             	     100	  29452113 ns/op
PASS
ok  	github.com/rueian/pg-perf-playground/multicolumns	223.744s
```