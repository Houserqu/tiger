import { RecoilRoot } from 'recoil';
import { BrowserRouter } from 'react-router-dom';
import Global from './component/Global';
import MyRouter from './routes';

function App() {
  return (
    <div className="App">
      <RecoilRoot>
        <BrowserRouter basename='/admin'>
          <Global>
            <MyRouter />
          </Global>
        </BrowserRouter>
      </RecoilRoot>
    </div>
  )
}

export default App
