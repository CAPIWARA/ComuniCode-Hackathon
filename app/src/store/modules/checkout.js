import { requestGraphQL } from '@/helpers/request';
import { CHECKOUT } from '@/store/types';

const actions = {
  [CHECKOUT]: async () => {
    await requestGraphQL(`
      mutation Checkout {
        checkout (value: "3000") {
          name
        }
      }
    `)
  }
};

export default { actions };
