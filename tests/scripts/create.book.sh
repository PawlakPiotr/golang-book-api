REQUEST_BODY='{"Title":"BOOK_TITLE","Author":{"FirstName":"Jan","LastName":"Kowalski"},"Category":"NEW_CATEGORY","Tags":["new_tag1","new_tag2"]}'

curl -X POST 'http://localhost:8080/api/v1/books/add' -H 'Content-Type: application/json' \
 -d $REQUEST_BODY