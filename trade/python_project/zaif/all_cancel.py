from zaifapi import *
import json

zaif = ZaifTradeApi(os.getenv("ZAIF_KEY"), os.getenv("ZAIF_SCT"))
active_order_list = zaif.active_orders()
print(active_order_list)
key_list = active_order_list.keys()
print(key_list)

for key in key_list:
    print(zaif.cancel_order(order_id=int(key)))
