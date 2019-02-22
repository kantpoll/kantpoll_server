## Before creating a campaign

The creator, before anything else, needs to have an idea of the number of voters who will participate in the campaign, what group categories he or she will create (for example, men and women), if there will be anonymous groups, if there will be observer nodes, what will be the start and end date of the campaign and if there will be an official website for the campaign.

If the campaign has a considerably large number of voters, the creator can use his or her organization's institutional email to validate them. Let's give an example. The creator of a campaign that involves all students, teachers and university employees could use as validation suffix "@someuniversity.edu" or "@election2019.someuniversity.edu".

There is the possibility of adding voters one by one. To do this, the creator of the campaign should ask voters to type "id" into the search bar (after they have created a vault) and copy the text that appears. Then they must send this text to the creator, who can register the voters and add them to groups. This can free the server from some unnecessary interactions.

## Creating a Campaign

Watch this [video](https://www.youtube.com/watch?v=JVGMO7cArDg).

## Managing a Campaign

To manage a campaign, click on the "Manage" button at the top of the campaign info page and follow the steps below:

1. (Optional) Insert other nodes in the network.
    Note: (Optional / Advanced) Add observers (for greater transparency).
2. Create group categories (example: men, women and anonymous).
3. Add a certain amount of groups of each category. Remember that each group can have 3, 27 or 81 voters.
    i. If necessary, add more groups. (Check the occupancy rate of the group category)
    ii. Leave the "Group administrator" field blank.
4. Insert a round.
    i. There may be up to five rounds.
    ii. The name of the round must contain up to 32 characters (letters, numbers and some punctuation marks).
    iii. When the round is created it is automatically set active.
5. Insert the candidates.
    Note: The candidate image link should start with "https" and end with an image file extension, for instance, jpg or png.
6. (Optional) Insert the parties.
7. Invite voters to join the campaign and choose a group (Or enter voters manually and choose groups for them - more practical).
    i. Do this by sharing the campaign link with them. This link can be found on the campaign info page just below the "Basic information".
    ii. Voters can choose their groups freely until the creator clicks "Stop group membership" on the "Groups" tab.
8. Voters vote for their candidates. They must wait for a *check* to appear on the line of their candidates.
9. After all voters have voted, the creator should go to the "Network" tab and click "Send" to send the votes.
    Note: If the votes have been sent correctly, the voters will see a *double check*.
10. After submitting the votes, the creator should go to the "Rounds" tab and then click on the "Close" button.
11. Click on "Count votes" in the "Network" tab. Voters should see a *double blue check*.
12. Go back to step 4 and create a new round, if applicable.

## How to add observers (Advanced)

Read the [README.txt](https://sourceforge.net/projects/kantcoin/files/)

## Test

If you want to simulate the insertion of a few hundred voters, you can do this:

1) Create a campaign.
2) Create a group category.
3) Create 10 groups of 81 voters.
4) Press CTRL + SHIFT + I and go to the console tab.
5) Paste this:
```
let results2 = "";
let results = "";
let seed2 = ethers.HDNode.mnemonicToSeed("tell zone wine demand model display noodle corn mention river shaft please");
let hdMaster2 = ethers.HDNode.fromSeed(seed2);
for(let i=0; i < 809; i++){
    let user = "user" + i;

    let key = hdMaster2.derivePath('m/0/' + i);
    let aux_signingkey = new ethers.SigningKey(key.privateKey);
    let pubkey = aux_signingkey.publicKey.substring(2);
    let address = aux_signingkey.address;
    
    results += user + ";" + address + ";" + pubkey + "\r\n"
    results2 += address + ";" + Math.floor((i) / 81) + ";" + (i % 81) + "\r\n"
}

insert_voters_textarea.value = results
```
6) On the admin page, in the GROUPS tab, press INSERT to insert voters. Wait a minute.
7) Go to the console tab and paste this:
```
add_voters_to_groups_textarea.value = results2
```
8) Press ADD to add voters to groups. Wait a few minutes.
