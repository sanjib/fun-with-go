### Windows PowerShell

Following is already set in user's path: `C:\Program Files\PostgreSQL\12\bin`

```
C:\> & { $env:PGUSER="postgres"; createuser -P -d gwp}
Enter password for new role: gwp
Enter it again: gwp
Password: postgres

C:\> createdb gwp
Password: postgres

// cd to repo

C:\lab\fun-with-go\chang\ch06> psql -U gwp -f setup.sql -d gwp
Password for user gwp: gwp
CREATE TABLE


```
