from zaifapi import *
import sys
from decimal import Decimal

argvs = sys.argv
argc = len(argvs)
print(str(argvs) + " count:" + str(argc))
if argc < 6 :
    exit(-1)
zaif = ZaifTradeApi(os.getenv("ZAIF_KEY"), os.getenv("ZAIF_SCT"))
print(zaif.trade(currency_pair=argvs[1],action=argvs[2],price=Decimal(argvs[3]),amount=Decimal(argvs[4]),comment=argvs[5]))

