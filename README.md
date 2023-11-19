<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a name="readme-top"></a>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![GPL-3.0 License][license-shield]][license-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/neutrino2211/NomadDB">
    <img src="docs/Nomad DB.png" alt="Logo" height="80">
  </a>

<h3 align="center">Nomad DB</h3>

  <p align="center">
    A decentralised P2P database format designed with an emphasis on ownership and anonymity which aims to support Individual and Communal instances through shared data with the hope of fostering a community of apps that store and share data on the network.
    <br />
    <a href="https://github.com/neutrino2211/NomadDB"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/neutrino2211/NomadDB">View Demo</a>
    ·
    <a href="https://github.com/neutrino2211/NomadDB/issues">Report Bug</a>
    ·
    <a href="https://github.com/neutrino2211/NomadDB/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

[![Nomad DB Architecture][product-screenshot]][product-screenshot]

A decentralised P2P database format designed with an emphasis on ownership and anonymity which aims to support Individual and Communal instances through shared data with the hope of fostering a community of apps that store and share data on the network.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

Install go>=1.20
* go
  - Go to [the golang download page](https://go.dev/doc/install) and follow the instructions
* Python
  - Go to [the python download page](https://www.python.org/downloads/) and follow the instructions


### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/neutrino2211/NomadDB.git
   ```
2. Install Go packages
   ```sh
   cd NomadDB && go get
   ```
3. Test run the program to make sure it is working
   ```sh
   go run .
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage
Only the most basic cluster instance exists and to run it, start a cluster.

```sh
go run . database
```

Then interact with the cluster on port 13564 with the python script in test-client/interact.py

* Write data
  
```sh
python interact.py -o <your_username> -d <content_to_store>
```

This should print some diagnostic messages and if it gets an "ok" response, it will print three things.
- ID of the file that stored the content within the cluster
- The content hash
- The length of the content hash [should be 64]

* Fetch data

```sh
python interact.py -o <your_username> -m fetch -r <the_content_hash>
```

This should print some diagnostic messages and if it gets an "ok" response, it will print the 2048 byte block content

* Delete data

```sh
python interact.py -o <your_username> -m delete -r <the_content_hash>
```

This should print some diagnostic messages and an empty byte string

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [X] Clusters
- [ ] Nodes of clusters
- [ ] Registries

See the [open issues](https://github.com/neutrino2211/NomadDB/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the GPL-3.0 License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Your Name - [@neutrino2211](https://twitter.com/neutrino2211) - tsowamainasara@gmail.com

Project Link: [https://github.com/neutrino2211/NomadDB](https://github.com/neutrino2211/NomadDB)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* [Best Read Me](https://github.com/othneildrew/Best-README-Template/tree/master)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/neutrino2211/NomadDB.svg?style=for-the-badge
[contributors-url]: https://github.com/neutrino2211/NomadDB/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/neutrino2211/NomadDB.svg?style=for-the-badge
[forks-url]: https://github.com/neutrino2211/NomadDB/network/members
[stars-shield]: https://img.shields.io/github/stars/neutrino2211/NomadDB.svg?style=for-the-badge
[stars-url]: https://github.com/neutrino2211/NomadDB/stargazers
[issues-shield]: https://img.shields.io/github/issues/neutrino2211/NomadDB.svg?style=for-the-badge
[issues-url]: https://github.com/neutrino2211/NomadDB/issues
[license-shield]: https://img.shields.io/github/license/neutrino2211/NomadDB.svg?style=for-the-badge
[license-url]: https://github.com/neutrino2211/NomadDB/blob/master/LICENSE
[arch-url]: https://github.com/neutrino2211/NomadDB/blob/master/docs/arch.png
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/linkedin_username
[product-screenshot]: docs/arch.jpg
[Next.js]: https://img.shields.io/badge/next.js-000000?style=for-the-badge&logo=nextdotjs&logoColor=white
[Next-url]: https://nextjs.org/
[React.js]: https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB
[React-url]: https://reactjs.org/
[Vue.js]: https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vuedotjs&logoColor=4FC08D
[Vue-url]: https://vuejs.org/
[Angular.io]: https://img.shields.io/badge/Angular-DD0031?style=for-the-badge&logo=angular&logoColor=white
[Angular-url]: https://angular.io/
[Svelte.dev]: https://img.shields.io/badge/Svelte-4A4A55?style=for-the-badge&logo=svelte&logoColor=FF3E00
[Svelte-url]: https://svelte.dev/
[Laravel.com]: https://img.shields.io/badge/Laravel-FF2D20?style=for-the-badge&logo=laravel&logoColor=white
[Laravel-url]: https://laravel.com
[Bootstrap.com]: https://img.shields.io/badge/Bootstrap-563D7C?style=for-the-badge&logo=bootstrap&logoColor=white
[Bootstrap-url]: https://getbootstrap.com
[JQuery.com]: https://img.shields.io/badge/jQuery-0769AD?style=for-the-badge&logo=jquery&logoColor=white
[JQuery-url]: https://jquery.com 