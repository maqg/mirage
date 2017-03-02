DROP VIEW IF EXISTS v_account;
CREATE VIEW v_account (Id,Name,State,Type,LastLogin,LastSync,CreateTime,Password,Email,PhoneNumber,Description) as SELECT ID,U_Name,U_State,U_Type,U_LastLogin,U_LastSync,U_CreateTime,U_Password,U_Email,U_PhoneNumber,U_Description FROM tb_account;

DROP VIEW IF EXISTS v_user;
CREATE VIEW v_user (Id,Name,State,Type,LastLogin,LastSync,CreateTime,Password,Email,PhoneNumber,Description) as SELECT ID,U_Name,U_State,U_Type,U_LastLogin,U_LastSync,U_CreateTime,U_Password,U_Email,U_PhoneNumber,U_Description FROM tb_user;
