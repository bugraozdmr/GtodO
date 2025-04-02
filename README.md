initializing db
sudo docker-compose -p gtodo up --build

used technologies
Docker
redis
postgres

packages
echo , gorm , uuid ,


rate limiting, asenkron
her ayın izlenme verileri

Batch işlemleri ile topluca veri gönderme (örneğin, her 1000 izlenme sonrası PostgreSQL'e kaydetme).

Tuzaklar (Traps) kullanarak, belirli koşullarda veritabanı güncellemesi yapma.

Veri arşivleme, sharding, rate limiting ve batch işlemleri gibi teknikler, veritabanı boyutunu kontrol altında tutmanıza yardımcı olur.