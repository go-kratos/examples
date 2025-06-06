export interface GrowCardItem {
  icon: string;
  title: string;
  value: number;
  total: number;
  color: string;
  action: string;
}

export const growCardList: GrowCardItem[] = [
  {
    title: '文章数',
    icon: 'post-count|svg',
    value: 2000,
    total: 120000,
    color: 'green',
    action: '周',
  },
  {
    title: '评论数',
    icon: 'comment-count|svg',
    value: 20000,
    total: 500000,
    color: 'blue',
    action: '周',
  },
  {
    title: '阅读量',
    icon: 'reading-count|svg',
    value: 8000,
    total: 120000,
    color: 'orange',
    action: '周',
  },
  {
    title: '访问数',
    icon: 'visit-count|svg',
    value: 5000,
    total: 50000,
    color: 'purple',
    action: '月',
  },
];
