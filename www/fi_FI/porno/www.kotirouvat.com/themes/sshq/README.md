# Setting Up This Theme

You must set up a config.toml when you install this theme on your system.
I've included "sample.config.toml" to help you with the setup.

The most important steps are:

1. Tell Hugo to use the theme.
2. Set the baseurl for your web site.
3. Change the title of the web site.
4. Add in your social media account information (or turn it off).
5. Set up Google Analytics for tracking (or turn it off).

If you copy the "sample.config.toml" file to your config.toml file,
search for "@@" to find settings that you must change to customize your
site. You will also need to update the theme parameters.

There are some values that you can set in your front matter. Doing so will
override the theme's defaults.

# Help

If you have questions, please visit http://discuss.gohugo.io. If you find
problems with this theme, please open an issue at http://github.io/quoha/sshq/.

# Variables

Please use variables in the configuration file and in your Markdown files
to control the theme.

## Config File Variables

The following variables are top level (meaning that they're not in a section).

1. title - This will display in the banner.

The following variables are in the "author" section.

1. name - If supplied, this will be displayed on every post.

The following variables are in the "Params.copyright" section. They are used to
populate the copyright footer on each page.

1. year - The string of years to display.
2. owner - The name of the owner to display.

The following variables are in the "Params.theme" section. For these variables,
the theme checks to see if they are set. It ignores the value. If 

1. slogan - This will display in the banner.
1. disqus_recent - If set, Recent Comments will display in the sidebar.
1. footer_powered_by - If set, display the "Powered By" text in the footer.
1. footer_designed_by - If set, display the "Designed By" text in the footer.
1. ga_tracking - If set, Javascript code to use Google Analytics will be added to every page.
1. twitter_follow - If set, a button to follow your Twitter account will display in the sidebar.
1. twitter_recent - If set, Recent Tweets will display in the sidebar.

Note that you must supply account information for Google Analytics, and Twitter if you wish to use those features.
For Disqus, you will need to supply the "shortName" used to reference your group.

## Social Account Setup

Discus, Google, and Twitter are supported. Generally speaking, if you don't supply
the account information then the theme will not include code to interact with the
site on your page.

### Disqus

Disqus is configured in the "Params.disqus" section. The following variables are used:

1. shortName - Your group name

If you're running Hugo on your local server then the theme will disable Disqus. There's
an issue when the hostname is "localhost", so we disable it.

### Google Search

Google Search is configured in the "Params.google" section. The following variables are used:

1. search_site - The domain name of your site. This isn't the URL, so it doesn't include "http".

### Google Analytics

Google Analytics is configured in the "Params.google.analytics" section. The following variables are used:

1. account - The name of your Google Analytics account.

### Twitter

Twitter is configured in the "Params.twitter" section. The following variables are used:

1. account - The name of your Twitter account.
1. ut_widget_id - Twitter's key that allows you to query for your recent tweets.

To display recent tweets, you must supply a Twitter Widget ID. If you don't then you will not
be able to display your recent tweets.

1. Open up twitter.com in your browser and sign in with your account.
1. Go to https://twitter.com/settings/widgets after signing in.
1. Click the "Create new" button.
1. Select the "User timeline" tab (it should be the default).
1. Enter your username.
1. Check "Exclude replies."
1. Uncheck "Auto-expand photos."
1. Click the "Create widget" button.
1. You should see a box with a lot of text in it. Below the box it should say "Copy and paste the code into the HTML of your site."
1. Copy that text into an editor (since it's hard to read in that tiny box).
1. Look for "data-widget-id" in the text. That will be followed by a very long number in quotes. That number is your widget ID.
1. If you forget the widget ID, you can edit the widget again.

## Front Matter Variables

You can also set the following variables in your front matter. This gives you control
in each post that you write.

1. no_comments = "true" - If set, then do not display comments for a post

# Logo

Most users won't be happy with the duck logo. That's okay because it is very easy to
change. The logo is stored in static/images/logo.png. The theme expects that it
will be a PNG file with a transparent background. The image should be 227 pixels
by 227 pixels, with a resolution of 72 dots per inch.

# Heritage

SSHQ uses the stylesheets, Javascript, and image files from the Octopress 2.0 theme by Brandon Mathis.
Octopress is a static site generator (SSG). I used on my website before switching over to Hugo.

For more information on Octopress, please visit http://octopress.org.

The theme uses Twitter-Post-Fetcher by Jason Mayes to grab recent tweets. The source is available at https://github.com/jasonmayes/Twitter-Post-Fetcher