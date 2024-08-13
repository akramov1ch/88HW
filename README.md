# Uy vazifasi: Go va Kafka bilan ishlash

## Maqsad
Ushbu vazifaning maqsadi - `Go` dasturlash tili yordamida `Kafka` uchun producer va consumer yaratishdir. Ushbu vazifa sizga real vaqtda ma'lumotlar oqimi bilan ishlashni va `Kafka`-ni `Go` dasturiga qanday integratsiya qilishni o‘rgatadi.


## Talablar
1. **Kafka-ni sozlash**:
    - `go-kafka-topic` nomli Kafka topic yarating.

2. **Kafka Producer:**:
    - `go-kafka-topic` ga `JSON` xabarlarini yuboradigan `Go` dasturida `Kafka` producer yarating.
    - `JSON` xabarlarida foydalanuvchi `ID`-si, message va timestamp belgisi bo'lishi kerak.

3. **Kafka Consumer:**:
    - `go-kafka-topic` dan xabarlarni o'qiydigan `Go` dasturida `Kafka` consumer yarating
    - O'qilgan xabarlarni konsolga log qiling.

4. **Partitsiyalash**:
    - Foydalanuvchi `ID`-ga asoslangan holda ma'lum partitsiyalarga xabarlarni yuboradigan producer ni o'zgartiring.
    - Consumer barcha partitsiyalardan xabarlarni o'qishini ta'minlang











