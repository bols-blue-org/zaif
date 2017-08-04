from zaifapi import *

zaif = ZaifTradeApi(os.getenv("ZAIF_KEY"), os.getenv("ZAIF_SCT"))
print(zaif.trade_history())

