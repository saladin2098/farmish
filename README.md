# Farmish

**Project Description:**
Fermerlar uchun qulaylashtirilgan, keng imkoniyatlarga ega kichkina API service.
Fermer o'zining fermasi haqida ma'lumot olishi jumladan, hayvonlari soni, kasal hayvonlar va och hayvonlar soni va ro'yxatini, hayvonlarning turiga qarab o'rtacha vaznigacha ko'rsatuvchi dashboard.
Hayvonlar uchun yemish oz qolganda 6 kun oldin, hayvonlar kasal bo'lganda va hayvonlarning qorni och qolganda, ovqatlantirish vaqti kelganligi haqida har 10s da bir ogohlatirish yuboradi. (Ogohlantirish oralig'ini qo'lda sozlab olish mumkin, ishlashi bilinib turishi uchun 10s qilib belgilandi.)
Ishlashini qulaylatirish uchun hayvonlar faqat 2, parranda va hayvon turlariga bo'lib olindi.

## Installation

1. Initialize a git repository and clone the project:
    ```sh
    git init
    git clone git@github.com:saladin2098/farmish.git
    ```
2. Create a database named `farmish` on port `5432`.
3. Update the `.env` file with the appropriate configuration.
   ```.env
   DB_HOST=localhost
   DB_USER=postgres
   DB_NAME=farmish
   DB_PASSWORD=pass
   DB_PORT=5432
   LOGPATH=logs/info.log
   ```

4. Use the following Makefile commands to manage the database migrations and set up the project:
    ```makefile
    # Set the database URL
    exp:
        export DBURL="postgres://mrbek:QodirovCoder@localhost:5432/farmish?sslmode=disable"

    # Run migrations
    mig-up:
        migrate -path migrations -database ${DBURL} -verbose up

    # Rollback migrations
    mig-down:
        migrate -path migrations -database ${DBURL} -verbose down

    # Create a new migration
    mig-create:
        migrate create -ext sql -dir migrations -seq create_table

    # Create an insert migration
    mig-insert:
        migrate create -ext sql -dir migrations -seq insert_table

    # Generate Swagger documentation
    swag:
        swag init -g api/handler.go -o api/docs

    # Clean up migrations (commented out by default)
    # mig-delete:
    #   rm -r db/migrations
    ```
5. Set the environment variable and run the project:
    ```sh
    make exp
    make mig-up
    go run main.go
    ```
6. Open the following URL to access the Swagger documentation:
    ```
    http://localhost:8080/api/swagger/index.html#/
    ```

## Features and Usages
1. Animal Create bo'limidan hayvon qo'sha oladi: (Hayvon qo'shayotganda default jsondagi ID datasi o'chirib tashlansin, hayvon qo'shayotganda animal_typega faqat hayvon yoki parranda kiritilsin. Agar hayvon sog'ligi true bo'lsa medication va conditionlarni kiritganligi inobatga olinmaydi).
2. Get all Animals bo'limida animal typei (parranda, hayvon yoki bo'sh), is_healthy (true, false yoki bo'sh), is_hungry (true, false yoki bo'sh) lariga qarab saralay oladi. (Filterlardan foydalanish ixtiyoriy va foydalanilmasa hamma hayvonlarni chiqaradi).
3. Feedingda schedule vaqti kelishidan 1 soat oldin feeding qilishni boshlay oladi va feeding schedule vaqti kelganda bildirishnoma yuboradi. Feeding schedule vaqti abetda ovqatlantirmasdan kechqurun ovqatlantirilsa oxirgi ovqatlantirilgan vaqti kechqurun deb hisoblanadi va hech qanday xatoliklar kelib chiqmaydi.
4. Agar scheduledan oldin ovqatlantirmoqchi bo'lsa schedulegacha ovqatlantirla olmaysiz degan error qaytadi va ovqatlantirilgan hayvonlarni ovqatlantirsa hayvonlar to'q deb error chiqaradi.
5. Hayvonlarni alohida animale_typega qarab ovqatlantira oladi. Bunda parrandalarni ovqatlantirsa provision (ombor) dagi ularning o'zining ovqatidan sarflanadi.
6. Ovqat sarflanish miqdori parranda bo'lsa uning qancha vaqt yashaganiga qarab hamma parrandalar uchun har xil hisonlanadi, agar u hayvon bo'lsa uning o'girligi ham inobatga olinadi. (ovqat iste'mol miqdori har kuni yangilanib turadi)
7. Agar hayvonlarga yoki parrandalarga bir birning ozuqasini bermoqchi bo'lsa xatolik beradi, hamma hayvonlar faqat o'zi yeyishi mumkin bo'lgan narsalar bilan ovqatlatira olinadi.
8. Agar medication qo'shayotgandau mavjud bo'lsa soni o'zgartiriladi va medications turiga qarab guruhlanib chiqariladi.
9. Agar provision (ombor)da 6 kunlikda kam ozuqa qolsa fermer bildirishnoma bilan ogohlantiriladi.
10. Ombordan hamma hayvonning iste'moliga qarab har safar har hil miqdorda ozuqa kamayadi. 

## Dependencies

- **Scheduling**: [github.com/go-co-op/gocron](https://github.com/go-co-op/gocron)
- **Swagger**: [github.com/swaggo/swag](https://github.com/swaggo/swag)
- **Database**:
    - [database/sql](https://golang.org/pkg/database/sql/)
    - [github.com/lib/pq](https://github.com/lib/pq)
- **Environment Variables**: [github.com/joho/godotenv](https://github.com/joho/godotenv)
- **API Framework**: [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- **Migrations**: [github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate)
****
## Acknowledgments

- Azizbek Qodirov
- Shamsiddin Okilov
- Feruza Mirjalilova

## Known Issues
1. Ba'zi testlarda kamchilik va xatoliklar bo'lishi mumkin;


## Special thanks to HUSAN MUSA