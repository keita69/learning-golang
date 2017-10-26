# learning-golang
Learn golang repository

```
@startuml
package utils/cs {
    class Main {
        + main()
        + select(num, args[])
    }

    interface MenuItem {
        + execute (num, args[])
    }

    class QuitMenuItem {
        + execute (num, args[])
    }

    Main .> MenuItem : menu item execute
    MenuItem -|> QuitMenuItem

}


@enduml
```

