# lab6-7
Lab materials for 6 + 7

## Lab 6:
`this is code you should run in a terminal`
> This is an aside, an optional or alternate explaination

### Setup
1. Setup an account on [heroku](https://heroku.com) (You should have done this over the weekend)
2. Make sure the heroku CLI tools are installed with `heroku version`, if not then install the [Heroku CLI tools](https://devcenter.heroku.com/articles/getting-started-with-go#set-up)
3. Make sure Go is installed on your machine using `go version`, if not [use this tutorial](https://golang.org/doc/install)
4. If you run into any issues, check [heroku's help center](https://devcenter.heroku.com/articles/getting-started-with-go#introduction)
5. Fork the lab repo at https://github.com/uw-info340b-sp2016/lab6-7 by going to the page and clicking the "fork" button. **Don't clone this repo! You won't be able to push to it, which is required to turn in your assignment!**
6. ~~Clone your newly forked repo from https://github.com/YOUR-USERNAME-HERE/lab6-7 to your machine using </br>`go get github.com/YOUR_GITHUB_USERNAME/lab6-7`~~ Skip this and use the below steps

    > This step is the same as: (replace $GOPATH with %GOPATH% on windows)
    >   1. `cd $GOPATH/src/github.com`
    >       * (If this step fails with "No such directory" then run   `cd $GOPATH` then `mkdir src`, `cd src`, `mkdir github.com`, and finally `cd github.com` )
    >   2. `mkdir YOUR_GITHUB_USERNAME`
    >   3. `cd YOUR_GITHUB_USERNAME`
    >   4. `git clone https://github.com/YOUR-USERNAME-HERE/lab6-7`
    >   5. ~~`cd lab6-7/cmd/lab7`~~
    >   6. ~~`go install`~~
    
If you run into any issues with these steps, especially if you get an issue with "go cannot find GOPATH" or something similar, follow [these steps on setting up Go](https://github.com/uw-info340b-sp2016/course/blob/master/guides/mac_env.md)

### Assignment
> For this assignment we will be using a new tool known as [PostgreSQL](http://www.postgresql.org/). This is different than SQLite in many ways, which, if you are interested, you can find out the [exact differences at this site](http://www.sql-workbench.net/dbms_comparison.html). Postgres is much better than SQLite in pretty much every way. It stores more datatypes, is well optimized, better checks on input, the list goes on. However, SQLite has a **much** smaller file size than Postgres so SQLite is best when you want one local DB running on a phone or simple desktop app. You only really want to use Postgres when you have a server running the DB, although there is nothing stopping you from running a local instance of Postgres.

#### Go and the website
1. Login to your heroku account and go to [your dashboard](https://dashboard.heroku.com/)
2. Create a new app ("+" in the top right corner, select "Create new app")
3. Set `app name = last-first-lab7` (*Actually* fill in your first and last name, lowercase only) and `runtime selection = United States`
4. Wait for the app to setup. This may take a minute or two.
5. While waiting, open your terminal and type `heroku login` then enter in your username + password to heroku. From now on don't close this terminal and run all "heroku" commands from here. If you do close it, you will need to `heroku login` again.
6. Open up your terminal and navigate to the project you cloned earlier. </br>`cd $GOPATH/src/github.com/YOUR_GITHUB_USERNAME/lab6-7`
7. Run `heroku git:remote -a last-first-lab7`
8. Push this code up to your new heroku app with `git push heroku master`
9. Open the app by running `heroku open` or by going to http://last-first-lab7.herokuapp.com/ The only error you should get is a missing database url, you can ignore that. **Note:** because the way heroku works, the free tier has a "dyno spin up time" if it is sleeping, and it may take up to 20 seconds for the page to load [See this page on sleeping for more info](https://devcenter.heroku.com/articles/dyno-sleeping)
10. Make a change to both the HTML + CSS, it doesn't have to be a huge change, but add some formatting or a new title. Whatever you feel like (**DON'T** delete the list element with ID='results', this is required for lab 7)
    
    > This next section will be basically the same steps as if you were pushing your code to GitHub, except it will be going to heroku instead

11. Run `git add .`, `git commit -m "put a message here"`, and finally  `git push heroku master`
12. Admire your new changes by refreshing the page you were on

#### Database

1. Go to https://postgres.heroku.com/databases (this can also be accessed by clicking on the 6 boxes in the top left of your dashboard)
2. Select "create database" in the top right. This will take a few minutes to prepare, but that's okay.
3. Click on the database name and wait for the page to load. Once it does, click "show" next to the URL attribute under "Connection Settings"
4. Copy this url. It should look something like postgres://z7fjdafjil:-dsFOSDJGc...
5. Go back to your dashboard for heroku, click on your app, click on setings, then click "Reveal config vars"
6. Add a new config var with the key `DATABASE_URL` and value as the prostgres url you copied. Then hit "add"
7. Refresh your heroku website and you should see the error message removed.

If you've gotten this far you are done for lab 6! I know it was a lot of setup, copying + pasting, and other boring things, but this is important as setting this stuff up takes quite a bit of time. Lab 7 will have you actually working with this database and application. For now though, there is nothing to turn in.

## Lab 7:

#### Update your fork with my changes
Run the following command *once* within your local reop. Once you do this you won't have to do it again (unless you work on lab machines)
`git remote add upstream https://github.com/uw-info340b-sp2016/lab6-7.git`

Then run the following command every time you want to get changes that I made in my repo
`git pull upstream master`

Depending on your text editor you may need to write a commit. In nano is is CTRL+O, CTRL+X and in vim this is :wq. There shouldn't be any merge conflicts, depending on what changes you made to the HTML.
