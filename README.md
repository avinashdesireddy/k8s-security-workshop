## PRO TALK - Kubernetes Security Workshop - CloudWorld 2022

This repository consists of the steps & manifests used to demo RBAC, OPA & Network Policies during CloudWorld 2022 PRO TALK Session.

### Role-Based Access Control
1. Create Users Blue, Green & Red
2. Using ClusterAdmin account- setup Blue, Green & Red Namespaces, create roles and rolebindings
   > From "rbac" directory
   ```
   kubectl apply -f 0.setup/setup-blue.yaml
   kubectl apply -f 0.setup/setup-green.yaml
   kubectl apply -f 0.setup/setup-red.yaml
   ```
3. Deploy applications using Blue user
   ```
   ## Success
   kubectl apply -f 1.deploy/blue-app.yaml

   ## Failure scenario --> User Blue trying to deploy app in "green" namespace
   kubectl apply -f 1.deploy/blue-app.yaml -n green
   ```
   > Verify if App is deployed and running successfully

4. Deploy applications using Green user
   ```
   ## Success
   kubectl apply -f 1.deploy/green-app.yaml

   ## Failure scenario --> User Green trying to deploy app in "red" namespace
   kubectl apply -f 1.deploy/green-app.yaml -n red
   ```
   > Verify if App is deployed and running successfully

5. Deploy applications using Red user
   ```
   ## Success
   kubectl apply -f 1.deploy/red-app.yaml

   ## Failure scenario --> User Red trying to deploy app in "default" namespace
   kubectl apply -f 1.deploy/red-app.yaml -n default
   ```
   > Verify if App is deployed and running successfully


### OPA/Gatekeeper

1. Install OPA
2. Create NodePort service for app "blue"
   ```
   kubectl apply -f nodeport-svc-blue.yaml
   ```
   > This will be SUCCESSFUL. 
   > Remove the service after validation
3. Create NodePort Policy
   ```
   cd node-port-policy

   ## Create Policy Template & Constraint
   kubectl apply -f 1.block-nodeport-template.yaml
   kubectl apply -f 2.block-nodeport-constraint.yaml
   ```
4. Redeploy NodePort service for app "blue"
   ```
   kubectl apply -f nodeport-svc-blue.yaml
   ```
   > This will fail.

5. Deploy the service with ClusterIP
   ```
   kubectl apply -f clusterip-svc-blue.yaml
   ```

### Network Policies
1. Re-deploy new version on Blue app
   ```
   kubectl apply -f deploy-blue-app-v2.yaml
   ```
2. Verify MySQL connectivity & API Connectivity by going into following URLs
   ```
   MySQL: http://<app-url>/mysql
   API: http://<app-url>/api
   ```
3. Create default deny all policy and refresh the page
   ```
   kubectl apply -f 1.default-deny-all.yaml
   ```
4. Open Firewall Rules
5. Create NetworkPolicy for MySQL Connectivity, and refresh the page
   ```
   kubectl apply -f 2.mysql-port-egress.yaml
   ```
6. Create NetworkPolicy for Green's API Connectivity
   ```
   kubectl apply -f 3.green-api-connect.yaml
   ```