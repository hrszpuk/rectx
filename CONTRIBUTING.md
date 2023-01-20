# Guidelines for Contributing
Thank you for considering contributing to ReCTx!
All kinds of contributions are appreciated from discussion to pull requests!

However, there are a few rules to follow that help keep rect organised and easier for other contributors to join in.
Although I use the word rules, these are more like guidelines. You don't have to follow everything here to the letter, 
but following along helps keep everything well organised and easy to understand.

## How to contribute
To begin contributing to ReCTx, you should first be aware of some systems in place.
Issues are used to keep track of what needs to be done.
Issues usually contain a description of the problem, potential solution, and a checklist for what needs to be done.
If you are looking to contribute, you should look into pre-existing issues and see if any take your fancy.

If you find an issue you want to help work on, you should respond in the comments that you would like to work on the 
solution and the issue will be assigned to you.
If you don't find any issues that you want to help with you can always submit your own issue!
By submitting your own issue you can make other contributors aware of issues in the code or enhancements that you think will make the project better.
When submitting your own issue, you should try to convey the problem and potential solution in the best way possible.
It's important the other contributors understand your ideas otherwise your issue could be closed.

After you have been assigned it's time to fork and branch!
You should commit changes to your own feature branch.
Once the checklist is complete, you can create a pull request.
The pull request needs to have a summary of the changes you've made (bullet point list is the best option here) and a reference to any related issues or pull requests.
Your code will be reviewed, and you may be asked to improve part of the code before the pull request is accepted.
Once accepted, your code will be merged into the master branch.

The master branch holds all the unstable changes to rectx.
Once a feature cycle is complete, the master branch pull request is created.
You may be requested to review the master branch code if you are a significant contributor.
Changes to the codebase may be made, but eventually a new release will be created and your code will be in a release version of rectx!

### Summary
To contribute to the project you will likely follow the following path:
1. Look for issues you want to help with or create your own issue.
   * Make sure to comment if you wish to help with an issue. 
   * When creating your own issue, follow the template and explain your ideas in a clear and concise way.
2. Forking and starting a branch.
   * Fork the project and create a feature branch!
   * After make commits to the feature branch, you can create a pull request.
3. When making a pull request you need to give a summary of the changes made and provide a reference to the issue(s) related.
4. Your pull request will be reviewed and changes may be requested.
   * If changes are requested it means something in your code is not right or could be improved.
   * You should continue to make changes until the review is approved and your code is ready to be merged.
5. Congratulations! Your code has been merged into the master branch!
   * You may have to wait until the next release to see your code in the final build.

### Branching strategy
ReCTx follows a common branching strategy where each feature/issue is given its own branch
These branches can have sub-branches but only if the issue is big enough to justify feature sub-branching.
Once the feature is ready, reviewed, and approved, the branch is merged into master.
Although it's in master, it's not actually in rectx yet!
That's because rectx uses a release tag system which means master can be unstable as people can download from the latest release.
