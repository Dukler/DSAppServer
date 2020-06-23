INSERT INTO users (email,username,password,token)
VALUES ('iarwain@hotmail.com.ar',
        'dukler',
        '$2a$08$SQQlz/GHkkOg4FQ7L/WJP.ChqeKGxIq2kZjOSSyHjG8oSFWm8YuAG',
        'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImlhcndhaW5AaG90bWFpbC5jb20uYXIifQ.gxMpRpH4dbcbijzyYqPO6jxCUD-dpDb9b9um_7pSAr4');
INSERT INTO app (name) VALUES ('mockApp');
INSERT INTO app (name) VALUES ('login');


DO $$
    DECLARE
        userID UUID;
        mockAppID UUID;
        loginAppID UUID;
        domainID UUID;
    BEGIN
        select id into userID from users where username = 'dukler';
        select id into mockAppID from app where name = 'mockApp';
        insert into domains (name,app_id) values ('localhost:3000',mockAppID);
        insert into domains (name,app_id) values ('duckstack.com',mockAppID);
        insert into domains (name,app_id) values ('duckstackui.firebaseapp.com',mockAppID);
        select id into loginAppID from app where name = 'login';
        select id into domainID from domains where name = 'localhost:3000';
        INSERT INTO users_app (user_id,app_id) VALUES (userID, mockAppID);
        INSERT INTO users_app (user_id,app_id) VALUES (userID, loginAppID);
        INSERT INTO users_domains (user_id,domain_id) VALUES (userID, domainID);
    END $$;



