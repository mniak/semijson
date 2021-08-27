### Definition

There are some strings that look like JSON but are not.

I'm calling these Semi-JSON, because they are partially.

### Purpose

I need a parser because I'm integrating with an undocumented platform. Through reverse engineering I found that they return data in this format.

### Examples:
```json
{
    key: "value"
}
```

```json
{
    date: new Date(2001, 12, 27)
}
```

It seems like the content is in fact piece of Javascript, but I'm not running this in my machine for security reasons,
And I don't even want to have a Javascript runtime installed, since I'm using Go.



