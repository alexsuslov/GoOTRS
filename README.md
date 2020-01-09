# GoOTRS

## Get by id

``` 
gotrs -get 744485  -AllArticles=true -Attachments=true DynamicFields=true | jq .

{
  "Ticket": [
    {
      "Age": 4485388,
      "PriorityID": "3",
      "ServiceID": "",
      "Type": "Default",
      "Responsible": "root@localhost",
...

```

## Update by id

```
cat test.json                                            
{
  "TicketID": 744485,
  "DynamicField": [{
    "Name": "VIP",
    "Value": "+"
  }]
}
```

```

cat test.json | gotrs -update=744485   

```
