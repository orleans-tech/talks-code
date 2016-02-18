class SearchProfile extends React.Component {
  constructor (props) {
    super(props);
    this.state = { searchQuery: '' }
  }

  onChange (event) {
    this.setState({
      searchQuery: event.target.value
    })
  }

  onSearchClick () {
    this.props.onSearch(this.state.searchQuery)
  }

  render () {
    return (
      <div style={styles.searchContainer}>
        <input type="text"
          placeholder='tsunammis (github)'
          onChange={this.onChange.bind(this)}
          style={styles.searchInput} />
        <button
          onClick={this.onSearchClick.bind(this)}
          style={styles.searchCta}>
          Rechercher
        </button>
      </div>
    )
  }
}


class Repositories extends React.Component {
  render () {
    const { data } = this.props
    return (
      <div>
        {data.map((repo) => <Repository key={repo.id} data={repo} />)}
      </div>
    )
  }
}

Repositories.propTypes = {
  data: React.PropTypes.array.isRequired
};


class Repository extends React.Component {
  render () {
    const { data } = this.props
    return (
      <div style={styles.repositoryContainer}>
        <a style={styles.repositoryLink} href={data.html_url} target='_blank'>
          {data.full_name}
        </a>
        <span style={styles.repositoryDesc}>
          {data.description}
        </span>
      </div>
    )
  }
}

Repository.propTypes = {
  data: React.PropTypes.object.isRequired
};


class ReposContainer extends React.Component {
  constructor (props) {
    super(props);
    this.state = {
      loaded: false,
      loading: false,
      notFound: false,
      repos: []
    }
  }

  componentDidMount () {
    this.load(this.props.username)
  }

  componentWillReceiveProps (nextProps) {
    this.load(nextProps.username)
  }

  load (username) {
    this.setState({ loading: true })

    fetch(`https://api.github.com/users/${username}/repos`)
      .then((resp) => {
        return resp.ok ? resp.json() : Promise.reject()
      })
      .then((repos_payload) => {
        console.log('load() success')
        this.setState({
          loaded: true,
          loading: false,
          notFound: false,
          repos: repos_payload
        })
      }, () => {
        console.log('load() error')
        this.setState({
          loaded: true,
          loading: false,
          notFound: true,
          repos: []
        })
      });
  }

  render () {
    if (!this.state.loaded && !this.state.loading) {
      return <div>{'Select an username'}</div>
    }

    if (this.state.loading) {
      return <div>{'loading ...'}</div>
    }

    if (this.state.notFound) {
      return <div>{'User not found'}</div>
    }

    return <Repositories data={this.state.repos} />
  }
}

ReposContainer.propTypes = {
  username: React.PropTypes.string.isRequired
};


class Page extends React.Component {
  constructor (props) {
    super(props)
    this.state = {
      searchQuery: 'tsunammis'
    }
  }

  onSearch (searchQuery) {
    console.log(`Page:onSearch(${searchQuery})`)
    this.setState({ searchQuery })
  }

  render () {
    return (
      <div style={styles.pageContainer}>
        <img style={styles.logo}
          src='./img/OrleansTech.png'
          alt='Orléans Tech Talks' />
        <SearchProfile onSearch={this.onSearch.bind(this)} />
        <ReposContainer username={this.state.searchQuery} />
      </div>
    )
  }
}

const styles = {
  pageContainer: {
    width: 700,
    margin: '0 auto'
  },
  repositoryContainer: {
    border: '1px solid #c8c8c8',
    backgroundColor: 'white',
    marginBottom: 6,
    padding: '3px 5px',
    borderRadius: 3,
    display: 'flex',
    flexDirection: 'row',
  },
  repositoryLink: {
    display: 'block',
    whiteSpace: 'nowrap'
  },
  repositoryDesc: {
    display: 'block',
    padding: '0 5px 0 10px',
    textOverflow: 'ellipsis',
    whiteSpace: 'nowrap',
    overflow: 'hidden',
    color: '#c8c8c8'
  },
  logo: {
    width: 250,
    display: 'block',
    margin: '20px auto 25px auto'
  },
  searchContainer: {
    display: 'flex',
    flexDirection: 'row',
    marginTop: 10,
    marginBottom: 25
  },
  searchInput: {
    height: 30,
    width: '100%',
    minWidth: 200,
    maxWidth: 500,
    padding: '0 10px 0 10px',
    margin: 0,
    fontSize: 14,
    border: '1px solid #c8c8c8',
    color: '#828282',
    borderRadius: 3
  },
  searchCta: {
    height: 30,
    display: 'block',
    border: '1px solid transparent',
    backgroundColor: 'transparent',
    textTransform: 'uppercase',
    padding: '0 5px 0 5px',
    margin: '0 0 0 5px',
    fontSize: 14,
    boxSizing: 'content-box',
    color: '#c8c8c8'
  }
}

ReactDOM.render(
  <Page />,
  document.getElementById('container')
);
