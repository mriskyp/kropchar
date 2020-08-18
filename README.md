# kropchar
kropchar is implement how many ways to generate random string and evaluate it


ignore main.go, as it only use as example

import to your code
```
go get "github.com/mriskyp/kropchar"
```

call generate random string on your code:
   in millisecond and microsecond
   ```
  GenerateRandomRuneString(lengthNumber)
  GenerateRandomByteString(lengthNumber)
  GenerateRandStringBytesMaskImprSrc(lengthNumber)
  GenerateRandStringBytesMaskImprSrcSB(lengthNumber)
  GenerateRandStringBytesMaskImprSrcUnsafe(lengthNumber)
  GenerateRandSeq(lengthNumber)
  ```
  
  in nano second
  ```
  GenerateRandomRemainderString(lengthNumber)
  GenerateRandomStringBytesMask(lengthNumber)
  ```
  
## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
