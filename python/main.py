from zaifapi import *

zaif = ZaifPublicApi()
print(zaif.last_price('btc_jpy'))

zaif = ZaifTradeApi(os.getenv("ZAIF_KEY"), os.getenv("ZAIF_SCT"))
print(zaif.get_info())
print(zaif.get_info2())
print(zaif.trade_history())
print(zaif.active_orders())

