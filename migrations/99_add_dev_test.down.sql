DO $$
    DECLARE
        userID UUID;
        mockAppID UUID;
        loginAppID UUID;
        domainID UUID;
    BEGIN
        select id into userID from users where username = 'dukler';
        select id into mockAppID from app where name = 'mockApp';
        select id into loginAppID from app where name = 'login';
        select id into domainID from domains where name = '192.168.1.2:3000';
        delete from users_app where user_id = userID;
        delete from users_domain where domain_id = domainID;
        delete from users where id = userID;
        delete from app where id = mockAppID or id = loginAppID;
        delete from domain where id = domainID;
    END $$;