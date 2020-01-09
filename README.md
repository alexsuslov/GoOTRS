# GoOTRS

## Use

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
