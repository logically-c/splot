FROM golang:1.13

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD echo '100.13470871572707\n112.49113728669916\n106.51378318623918\n117.43779244036739\n110.66632474928402\n97.6549208455096\n107.24458360792656\n90.59000144578327\n107.71527424750806\n110.71429518872388\n113.38015808818254\n112.15287076612272\n109.63439828838985\n100.13470871572707\n112.49113728669916\n106.51378318623918\n117.43779244036739\n110.66632474928402\n97.6549208455096\n107.24458360792656\n90.59000144578327\n107.71527424750806\n110.71429518872388\n113.38015808818254\n112.15287076612272\n109.63439828838985\n100.13470871572707\n112.49113728669916\n106.51378318623918\n117.43779244036739\n110.66632474928402\n97.6549208455096\n107.24458360792656\n90.59000144578327\n107.71527424750806\n110.71429518872388\n113.38015808818254\n112.15287076612272\n109.63439828838985' | app -m '*' 

# CMD echo '99\n0\n23\n54\n16\n34\n12\n16\n73\n82\n53\n32\n78\n11\n1' | app -m="-"