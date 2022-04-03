INSERT INTO user (
    email, 
    first_name, 
    last_name, 
    password, 
    phone, 
    is_alive, 
    role) 
VALUES (
    " + oneUser.Email + " ," 
    + oneUser.Firstname + " ," 
    + oneUser.Lastname + " ," 
    + oneUser.Password + " ," 
    + oneUser.Phone + " ," 
    + strconv.FormatBool(oneUser.IsAlive) + " ,(
        SELECT id FROM role 
        WHERE role = '" + oneUser.Role + "'
        );"