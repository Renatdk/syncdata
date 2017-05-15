#О загрузке данных по торговым рынкам

Данные беруться с сайта
https://blog.quandl.com/api-for-commodity-data

В файле configs/database.json
находятся данные для подключения к базе данных

configs/config.json
настройки по загрузки данных с API сервера

в поле данные
 [
    "ICE Brent Crude Oil Futures",
    "BrentOil",
    "CHRIS/ICE_B1",
    "4"
]

"4" - это позиция элемента в ыгружаемая с массива элементов