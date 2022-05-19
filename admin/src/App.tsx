import { RecoilRoot } from 'recoil';
import { unstable_HistoryRouter as HistoryRouter } from 'react-router-dom';
import Global from './component/Global';
import MyRouter from './routes';
import history from './history'

function App() {
  return (
    <div className="App">
      <RecoilRoot>
        <HistoryRouter history={history} basename='/admin'>
          <Global>
            <MyRouter />
          </Global>
        </HistoryRouter>
      </RecoilRoot>
    </div>
  )
}

export default App
