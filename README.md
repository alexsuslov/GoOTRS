# GoOTRS
## Help
```
./gotrs -help
Usage of ./gotrs:
  -AllArticles
        get AllArticles from tiket
  -Attachments
        get Attachments from tiket
  -DynamicFields
        get  DynamicFields from tiket
  -config string
        gotrs config env (default ".env")
  -debugger
        enable debugger
  -get string
        get tiket from OTRS
  -update string
        update tiket in OTRS

```

## Get by id

``` 
gotrs -get 744485  -AllArticles=true -Attachments=true -DynamicFields=true | jq .

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

### Error Message

```json
{
  "Error":{
    "ErrorCode":"TicketUpdate.InvalidParameter",
    "ErrorMessage":"TicketUpdate: DynamicField->Name parameter is invalid!"
  }
}
```
