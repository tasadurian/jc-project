# jc-project

1.   Hash   and   Encode   a   Password   String
Provide   the   ability   to   take   a   password   as   a   string   and   return   a   Base64   encoded   string   of   the
password   that   has   been   hashed   with   SHA512   as   the   hashing   algorithm.
For   example,   if   given   the   string    “angryMonkey”    the   expected   return   value   is 
`“ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZP ZklJz0Fd7su2A+gf7Q==”`

2.   Hash   and   Encode   Passwords   over   HTTP
Change   your   program   so   that   when   launched   your   code   starts   and   listens   for   HTTP   requests   on a   provided   port.   Listen   for   a    POST    to    `/hash`    with   a   body   of   the   form `password=<the_provided_password>` .   The   response   should   be   the   base64   encoded string   of   the   SHA512   hash   of   the   provided   password.   The   server   should   not   respond immediately,   it   should   leave   the   socket   open   for   5   seconds   before   responding.
For  example:   `curl   —data   “password=angryMonkey” http://localhost:8080/hash`    should   return `“ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZP ZklJz0Fd7su2A+gf7Q==”`    after   approximately   5   seconds.
Note:   Your   software   should   be   able   to   process   multiple   connections   simultaneously. 

3.   Graceful   Shutdown
Provide   support   for   a   “graceful   shutdown   request”.   If   a   request   is   made   to    `/shutdown`    all remaining   requests   should   be   allowed   to   complete,   no   additional   requests   should   be   allowed   to be   processed,   and   the   server   should   stop.

4.   Hash   End-Point   Returns   Identifier
Modify   the    `/hash`    end-point   so   that   it   returns   an   identifier   that   can   be   used   later   to   retrieve   a hashed   password   value.   The   request   should   be   initiated   identically   to   step   2.   The   identifier should   return   immediately   but   the   password   should   not   be   hashed   for   5   seconds.
For  example:   `curl   —data   “password=angryMonkey” http://localhost:8080/hash`    should   return    42    immediately.

5.   GET   a   Hashed   Password
Provide   the   ability   to   retrieve   an   already   hashed   password   with   a  G   ET    to    `/hash/{id}`
To   continue   our   example   from   the   last   step,   5   seconds   after   that  c   url    command   is   issued   you should  be  able  to   curl   `http://localhost:8080/hash/42`   and  get  back  the  value  of `“ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZP ZklJz0Fd7su2A+gf7Q==”`

6.   Statistics   End-Point
Provide   a   statistics   endpoint   to   get   basic   information   about   your   password   hashes.
A    GET    request   to    `/stats`    should   return   a   JSON   object   with   2   key/value   pairs.   The  “   total” key   should   have   a   value   for   the   count   of   password   hash   requests   made   to   the   server   so   far.   The “average”    key   should   have   a   value   for   the   average   time   it   has   taken   to   process   all   of   those requests   in   milliseconds.
For  example:   `curl   http://localhost:8080/stats`   should  return  something  like: `{“total”:   1,   “average”:   123}`