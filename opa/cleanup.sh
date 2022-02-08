## Remove existing deployments
kubectl delete deployment blue -n blue
kubectl delete deployment green -n green
kubectl delete deployment red -n red

kubectl delete service blue-svc -n blue
kubectl delete service green-svc -n green
kubectl delete service red-svc -n red
