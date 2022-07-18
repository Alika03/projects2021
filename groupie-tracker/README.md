# groupie-tracker
# Description
Groupie Trackers consists on receiving a given API and manipulates the data contained in it, in order to create a site, displaying the information.

    It will be given an API, that consists in four parts:

        The first one, artists, containing information about some bands and artists like their name(s), image, in which year they began their activity, the date of their first album and the members.

        The second one, locations, consists in their last and/or upcoming concert locations.

        The third one, dates, consists in their last and/or upcoming concert dates.

        And the last one, relation, does the link between all the other parts, artists, dates and locations.

Groupie tracker search bar consists of creating a functional program that searches, inside your website, for a specific text input.

    The program handles these search cases :
        - artist/band name
        - members
        - locations
        - first album date
        - creation date
    The program handles case sensitive.
    The search bar has typing suggestions as you write.
        The search bar identifies and displays in each suggestion the individual type of the search cases. (ex: Freddie Mercury -> member)
        For example if you start writing "phil" it should appear as suggestions Phil Collins - member and Phil Collins - artist/band. This is just an example of a display.



Groupie Tracker Filters consists on letting the user filter the artists/bands that will be shown.

    Project incorporate these four filters:
       - Filter by creation date
       - Filter by first album date
       - Filter by number of members
       - Filter by locations of concerts

    It used two types of filters:
       - A range filter
       - A check box filter


Groupie Tracker Geolocalization consists on mapping the different concerts locations of a certain artist/band given by the Client.

Groupie tracker visualizations consists of manipulating the data coming from the API and displaying it in the most presentable way possible to you. The Schneiderman's 8 Golden Rules of Interface Design is followed:

   - Strive for consistency
   - Enable frequent users to use shortcuts
   - Offer informative feedback
   - Design dialogue to yield closure
   - Offer simple error handling
   - Permit easy reversal of actions
   - Support internal locus of control
   - Reduce short-term memory load

# Instruments

- The backend is written in Go.
- Only the standard Go packages are used.
- JSON files and format.
- Go routines.
- HTML/CSS/JS (The basics of human-computer interface)

# Image
1. ![alt text](https://github.com/Alika03/groupie-tracker/blob/272e63bfad3c6c264982eda5285335b6ed9e30c6/image/Screenshot%20from%202021-11-15%2016-51-32.png?raw=true)
2. ![alt text](https://github.com/Alika03/groupie-tracker/blob/272e63bfad3c6c264982eda5285335b6ed9e30c6/image/Screenshot%20from%202021-11-15%2017-08-05.png?raw=true)
3. ![alt text](https://github.com/Alika03/groupie-tracker/blob/272e63bfad3c6c264982eda5285335b6ed9e30c6/image/Screenshot%20from%202021-11-15%2017-08-32.png?raw=true)

**Usage to run :**
1. Run on the Terminal:`go run .`
2. Open **localhost:8080** on your browser
The authors of this project: @Alika03

