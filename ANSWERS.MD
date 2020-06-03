* How long did it take to finish the test?

It took me more than expected, I've been working in this test for around 3h

* If you had more time, how could you improve your solution?

I really would like to have work on make it https available, but I didn't want to put the self signed certs in the repository, and the solutions that I've found required more investigation. Since the challenge asked for simplicity over complexity I decided to go without https.

+ What changes would you do to your solution to be deployed in a production environment?

This solution needs some work in order to be production-ready:

1. SSL configuration should be a must if it's going to be production
2. Adding to `deployment.yaml` configuration about resource limiting should be done as well.
3. Adding an endpoint like `/ping` in order to know if the application is alive or not
4. Configuring readinessProbe with the endpoint commented in point 3
5. Adding some monitoring to get metrics and find the best point to autoscaling cluster (CPU, memory, etc)

* Why did you choose this language over others?

I'm learning Go, so I try to do everything I can in Golang, so I can sharp my skills :)

* What was the most challenging part and why?

Deciding which is the best method to deploy the solution. I decided to go with Kustomize because I've been working with it and I found it easy once you learn the basics.