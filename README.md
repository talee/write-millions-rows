# write-millions-rows
Create a CSV with millions of rows.

## Optional
Modify MySQL config /usr/local/etc/my.cnf to include

    secure_file_priv = ''
    
Or it's more secure variant of an actual directory where your output file will be. Otherwise by default MySQL prevents loading data into tables by file.


Then run this query

    LOAD DATA INFILE '/full/path/write-millions-rows/output_table_content.csv' INTO TABLE your_table_name;
