import IconVolunteering from '@/assets/images/Help-Volunteering.png';
import IconMoney from '@/assets/images/Help-Money.png';
import IconOther from '@/assets/images/Help-Other.png';

export default [
  {
    image: IconVolunteering,
    title: 'Trabalho Voluntário',
    description: 'Trabalho voluntário inclui todas as idades, diversos grupos religiosos e étnicos. Indivíduos com variados níveis de habilidade e recursos econômicos podem fornecer serviços úteis.',
    button: 'Faça parte',
    link: '/sign-on'
  },
  {
    image: IconMoney,
    title: 'Apadrinhe uma criança',
    description: 'Associados-colaboradores contribuem mensalmente por meio de quantias e em períodos por eles estipulados',
    button: 'Doe agora',
    link: '/checkout'
  },
  {
    image: IconOther,
    title: 'Diversas ações',
    description: 'Aceitamos roupas, materiais escolares, alimentos, brinquedos e calçados.As doações podem ser entregues diretamente no Projeto Madre Cabrini.',
    button: 'Contato',
    link: '/#contact'
  }
];
