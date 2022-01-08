

1.) bitcoin-cli stop

2.) Find chainstate folder of bitcoin; if different from default. Use -db /chaincopy/

3.) run: utxodump
                   
4.) Afters output to utxodump.csv (program dir) run:

5.) utxoparse greater 10000000
                     ^
              (lesser or equal) (comparatory number of satoshis)


## Dumps CSV like this
count,txid,vout,amount,type,address
